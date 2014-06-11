#! /usr/bin/perl

# Copyright 2012 The go-gl Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This script takes a file containing glew.h code as argument and spits it out as Go source.
# The contents should contain the typedefs and functions as blocks separated by a single blank line:
#
# typedef void (GLAPIENTRY * PFNGLBINDSAMPLERPROC) (GLuint unit, GLuint sampler);
# ...
#
# #define glBindSampler GLEW_GET_FUN(__glewBindSampler)
# ...
#
# The conversion is not 100% compilable so some manual fixing may be required.

use strict;
use warnings;

my @typedefs;
my @functions;

while (<>) {
	chomp;
  	last if(/^$/);
	push @typedefs, $_;
}

while (<>) {
	chomp;
	next if(/^$/);
	my ($funcName) = /^#define\s+(.+)\s.+$/;
	push @functions, $funcName;
}

die "Amount of typedefs does not match function count" if (@typedefs != @functions);

foreach my $i (0..$#typedefs) {
	$_ = $typedefs[$i];
	my ($return, $params) = /^
		typedef\s+
		([^\(]+)
		\s+
		\(.*?\)
		\s*
		\(([^\)]+)\)
	/x;

	$params =~ s/type/typ/;
	$params =~ s/range/rang/;
	$params =~ s/map/mp/;
	$params =~ s/const\s+//;

	my @goParams = split /\s*,\s*/, $params;
	foreach my $param (@goParams) {
		$param = glToGoType($param);
	}
	my $goParams = join ', ', @goParams;
	my @glParams = split /\s*,\s*/, $params;
	foreach my $param (@glParams) {
		$param =~ s/(\S+)\s+(\S+)/C.$1($2)/;

		$param =~ s/C.GLboolean/glBool/;
		$param =~ s/C.GLvoid\*/glPointer/;
		$param =~ s/(C.[^\*]+)\*\(([^\)])\)/(*$1)(&$2\[0\])/;
	}
	my $glParams = join ', ', @glParams;

	my $glFuncName = $functions[$i];
	my ($goFuncName) = $glFuncName =~ /^gl(.+?D?)[A-Z]*$/;

	$return = $return eq "void" ? "" : glToGoType("$return ");

	print
"
func $goFuncName\($goParams\) $return\{
	C.$glFuncName\($glParams\)
\}
";
}

sub glToGoType {
	my $type = $_[0];
	$type =~ s/(\S+)\s+(\S+)/$2 $1/;
	$type =~ s/GLbitfield/uint32/;
	$type =~ s/GLboolean/bool/;
	$type =~ s/GLbyte/int8/;
	$type =~ s/GLclampd/float64/;
	$type =~ s/GLclampf/float32/;
	$type =~ s/GLdouble/float64/;
	$type =~ s/GLfloat/float32/;
	$type =~ s/GLint/int/;
	$type =~ s/GLshort/int16/;
	$type =~ s/GLsizei/int/;
	$type =~ s/GLubyte/uint8/;
	$type =~ s/GLuint/uint/;
	$type =~ s/GLulong/uint64/;
	$type =~ s/GLushort/uint16/;
	$type =~ s/intptr/int/;
	$type =~ s/sizeptr/int/;

	$type =~ s/GLvoid\*/interface\{\}/;
	$type =~ s/GLchar\*/string/;
	$type =~ s/(\S+)\*/\[\]$1/;

	return $type;
}
