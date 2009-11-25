#!/usr/bin/perl

sub escape_keyword($)
{

    my @keywords=(
    "break", "default", "func", "interface", "select",
    "case", "defer", "go", "map", "struct",
    "chan", "else", "goto", "package", "switch",
    "const", "fallthrough", "if", "range", "type",
    "continue", "for", "import", "return", "var"
    );

    my $name=shift;

    for(@keywords)
    {
        if($_ eq $name)
        {
            return $name."_";
        }
    }

    return $name;

}

sub trim($)
{
	my $string = shift;
	$string =~ s/^\s+//;
	$string =~ s/\s+$//;
	return $string;
}

sub read_type($)
{

    my $arg = shift;

    my $ptr="";

    while($arg =~ s/\*/ / || $arg =~ s/\[\d*\]/ / )
    {
        $ptr.="*";
    }

    $arg =~ s/__attribute__\s*\(\(.*\)\)//g;

    $arg =~ s/\bconst\b/ /g;
    $arg =~ s/\bvoid\b/ /g;
    $arg =~ s/\bextern\b/ /g;

    $arg =~ s/\s+/ /g;
    $arg = trim($arg);

    if($arg eq "")
    {
        return {name => "",type => "",ptr => ""};
    }

    if( $arg =~ /^(\w+)( \w+)?$/)
    {
        my $name=$2?escape_keyword(trim($2)):undef;
        my $type=$1;

        # convert void* -> unsafe.Pointer
        if($ptr =~ /[*]/ && ($type eq "void" || $type eq "GLvoid"))
        {
            $ptr =~ s/[*]//;
            $type = "unsafe.Pointer";
        }

        # convert char* -> CString 
        if($ptr eq "*" && ($type eq "GLchar" || $type eq "char"))
        {
            $ptr =~ s/[*]//;
            $type = "string";
        }

        return {name => $name,type => $type,ptr => $ptr};
    }
    else
    {
#        die("error reading type '$arg'");
        print STDERR "error reading type '$arg'";
        return null;
    }

}

sub print_go_type($)
{
    my $type=shift;

    return trim("$type->{name} $type->{ptr}$type->{type}");

}

sub print_c_cast($)
{
    my $type=shift;

    my $cast;
    my $name=$type->{name};


    if($type->{type} eq "unsafe.Pointer")
    {
        $cast="";
    }
    elsif($type->{type} eq "string")
    {
        return "($type->{ptr}*_C_GLchar)(($type->{ptr}C.CString)($type->{name}))";
    }
    else
    {
        $cast="$type->{ptr}C.$type->{type}";
    }

    if($cast eq "")
    {
        return $name;
    }
    elsif($cast =~ /[*]/)
    {
        return "($cast)($name)"
    }
    else
    {
        return "$cast($name)"
    }

}

sub function_name($)
{

    my $name=shift;

#TODO remove ME
    $name =~ s/^glew//;
    $name =~ s/^gl//;
#

    $name =~ s/^(\w)/\u$1/;
    $name=escape_keyword($name);

    return $name;

}

my $package=$ARGV[0];
$package =~ s/\..*$//;

print "package $package\n\n";

open FD, "$ARGV[0]" or die "cat open";

while(my $line=<FD>)
{
    if($line =~ /\w/)
    {
        print "// $line";
    }
}

print "import \"C\"\n";
print "import \"unsafe\"\n";
print "\n";


close FD;

open FD, "cpp $ARGV[0]|" or die "cat open";

my $text = "";

while(my $line=<FD>)
{

    if(! ($line=~/^\s*#/))
    {
        $text.=$line;
    }
}

foreach my $dec (split(';',$text))
{


    $dec =~ s/\s+/ /;
    $dec = trim($dec);

    if($dec =~ /\s*typedef /)
    {
        if($dec =~ /(\w+$)/)
        {
            my $name = $1;
            if(! ($name =~ /void$/))
            {
                print "type $name C.$name;\n";
            }
        }
    }
    elsif($dec =~ /(.*) (\w+)\s*\((.*)\)/)
    {

        my $name=trim($2);

        my $return=trim($1);
        my @params=split(',',$3);
        my @args;

        my $return_type=read_type($return);
        $return=print_go_type($return_type);

        
        if($name =~ /^glu/)
        {
            next;
        }

        print "//".$dec."\n";


        my $i=0;

        foreach my $arg (@params)
        {

            $arg=trim($arg);

            if($arg ne "void" && $arg =~ /\w/)
            {

                my $type=read_type($arg);

                if(!$type->{name})
                {
                    $type->{name}="arg$i";
                }

                push(@args,$type);

                $i++;

            }

        }

        print "func ".function_name($name)."(";

        my $first=1;

        foreach my $arg (@args)
        {
            if($first)
            {
                $first=0;
            }
            else
            {
                print ", ";
            }

            print print_go_type($arg);
        }

        print ") $return\n";
        print "{\n";

        if($return ne "")
        {
            print "\treturn ".(($return =~ /[*]/)?"($return)":$return)."(";
        }
        else
        {   
            print "\t";
        }

        print "C.$name(";

        my $first=1;

        foreach my $arg (@args)
        {
            if($first)
            {
                $first=0;
            }
            else
            {
                print ", ";
            }

            print print_c_cast($arg);


        }

        print ")";

        if($return ne "")
        {
            print ")"
        }

        print ";\n";

        print "}\n\n";
    }
    else
    {
        #print $dec."\n";
    }

}

close FD;

