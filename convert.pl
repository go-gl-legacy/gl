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

    if( $arg =~ /(.*[^\w])(\w+$)/)
    {
        my $name=escape_keyword(trim($2));
        my $type=$1;
        my $ptr="";
        $type =~ s/const //g;

        if($type =~ /\*/)
        {
            $ptr = $type;
            $ptr =~ s/[^*]//g;
        }

        $type =~ s/[ *]+//g;

        if($ptr =~ /[*]/ && ($type eq "void" || $type eq "GLvoid"))
        {
            $ptr =~ s/[*]//;
            $type = "unsafe.Pointer";
        }



        return {name => $name,type => $type,ptr => $ptr};
    }
    else
    {
        die("error reading type '$arg'");
    }

}

sub function_name($)
{

    my $name=shift;

    $name =~ s/^gl//;

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
    elsif($dec =~ /(.*) (\w+) \((.*)\)/)
    {

        print "//".$dec."\n";

        my $name=trim($2);
        my $return=trim($1);
        my @params=split(',',$3);
        my @args;

        $return =~ s/const//g;
        $return =~ s/void//g;

        if($return =~ /[*]/)
        {
            $return =~ s/[*]+//g;
            $return = "*".trim($return);
        }

        foreach my $arg (@params)
        {

            $arg=trim($arg);

            if($arg ne "void" && $arg =~ /\w/)
            {
                push(@args,read_type($arg));
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

            print "$arg->{name} $arg->{ptr}$arg->{type}";
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

            if($arg->{type} eq "unsafe.Pointer")
            {
                print "$arg->{name}";
            }
            elsif($arg->{ptr} eq "")
            {
                print "C.$arg->{type}($arg->{name})";
            }
            else
            {
                print "($arg->{ptr}C.$arg->{type})($arg->{name})";
            }

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

