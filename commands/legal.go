// Copyright Datajin Technologies, Inc. 2015,2016. All rights reserved.
// Use of this source code is governed by an Artistic-2
// license that can be found in the LICENSE file.

package commands

import "github.com/codegangsta/cli"

const artisticShort = `
Copyright Datajin Technologies, Inc. 2015,2016. All rights reserved.
This software is released under the Artistic License 2.0.
`
const artisticFull = `
               The Artistic License 2.0

           Copyright (c) 2000-2006, The Perl Foundation.

     Everyone is permitted to copy and distribute verbatim copies
      of this license document, but changing it is not allowed.

Preamble

This license establishes the terms under which a given free software
Package may be copied, modified, distributed, and/or redistributed.
The intent is that the Copyright Holder maintains some artistic
control over the development of that Package while still keeping the
Package available as open source and free software.

You are always permitted to make arrangements wholly outside of this
license directly with the Copyright Holder of a given Package.  If the
terms of this license do not permit the full use that you propose to
make of the Package, you should contact the Copyright Holder and seek
a different licensing arrangement.

Definitions

    "Copyright Holder" means the individual(s) or organization(s)
    named in the copyright notice for the entire Package.

    "Contributor" means any party that has contributed code or other
    material to the Package, in accordance with the Copyright Holder's
    procedures.

    "You" and "your" means any person who would like to copy,
    distribute, or modify the Package.

    "Package" means the collection of files distributed by the
    Copyright Holder, and derivatives of that collection and/or of
    those files. A given Package may consist of either the Standard
    Version, or a Modified Version.

    "Distribute" means providing a copy of the Package or making it
    accessible to anyone else, or in the case of a company or
    organization, to others outside of your company or organization.

    "Distributor Fee" means any fee that you charge for Distributing
    this Package or providing support for this Package to another
    party.  It does not mean licensing fees.

    "Standard Version" refers to the Package if it has not been
    modified, or has been modified only in ways explicitly requested
    by the Copyright Holder.

    "Modified Version" means the Package, if it has been changed, and
    such changes were not explicitly requested by the Copyright
    Holder.

    "Original License" means this Artistic License as Distributed with
    the Standard Version of the Package, in its current version or as
    it may be modified by The Perl Foundation in the future.

    "Source" form means the source code, documentation source, and
    configuration files for the Package.

    "Compiled" form means the compiled bytecode, object code, binary,
    or any other form resulting from mechanical transformation or
    translation of the Source form.


Permission for Use and Modification Without Distribution

(1)  You are permitted to use the Standard Version and create and use
Modified Versions for any purpose without restriction, provided that
you do not Distribute the Modified Version.


Permissions for Redistribution of the Standard Version

(2)  You may Distribute verbatim copies of the Source form of the
Standard Version of this Package in any medium without restriction,
either gratis or for a Distributor Fee, provided that you duplicate
all of the original copyright notices and associated disclaimers.  At
your discretion, such verbatim copies may or may not include a
Compiled form of the Package.

(3)  You may apply any bug fixes, portability changes, and other
modifications made available from the Copyright Holder.  The resulting
Package will still be considered the Standard Version, and as such
will be subject to the Original License.


Distribution of Modified Versions of the Package as Source

(4)  You may Distribute your Modified Version as Source (either gratis
or for a Distributor Fee, and with or without a Compiled form of the
Modified Version) provided that you clearly document how it differs
from the Standard Version, including, but not limited to, documenting
any non-standard features, executables, or modules, and provided that
you do at least ONE of the following:

    (a)  make the Modified Version available to the Copyright Holder
    of the Standard Version, under the Original License, so that the
    Copyright Holder may include your modifications in the Standard
    Version.

    (b)  ensure that installation of your Modified Version does not
    prevent the user installing or running the Standard Version. In
    addition, the Modified Version must bear a name that is different
    from the name of the Standard Version.

    (c)  allow anyone who receives a copy of the Modified Version to
    make the Source form of the Modified Version available to others
    under

    (i)  the Original License or

    (ii)  a license that permits the licensee to freely copy,
    modify and redistribute the Modified Version using the same
    licensing terms that apply to the copy that the licensee
    received, and requires that the Source form of the Modified
    Version, and of any works derived from it, be made freely
    available in that license fees are prohibited but Distributor
    Fees are allowed.


Distribution of Compiled Forms of the Standard Version
or Modified Versions without the Source

(5)  You may Distribute Compiled forms of the Standard Version without
the Source, provided that you include complete instructions on how to
get the Source of the Standard Version.  Such instructions must be
valid at the time of your distribution.  If these instructions, at any
time while you are carrying out such distribution, become invalid, you
must provide new instructions on demand or cease further distribution.
If you provide valid instructions or cease distribution within thirty
days after you become aware that the instructions are invalid, then
you do not forfeit any of your rights under this license.

(6)  You may Distribute a Modified Version in Compiled form without
the Source, provided that you comply with Section 4 with respect to
the Source of the Modified Version.


Aggregating or Linking the Package

(7)  You may aggregate the Package (either the Standard Version or
Modified Version) with other packages and Distribute the resulting
aggregation provided that you do not charge a licensing fee for the
Package.  Distributor Fees are permitted, and licensing fees for other
components in the aggregation are permitted. The terms of this license
apply to the use and Distribution of the Standard or Modified Versions
as included in the aggregation.

(8) You are permitted to link Modified and Standard Versions with
other works, to embed the Package in a larger work of your own, or to
build stand-alone binary or bytecode versions of applications that
include the Package, and Distribute the result without restriction,
provided the result does not expose a direct interface to the Package.


Items That are Not Considered Part of a Modified Version

(9) Works (including, but not limited to, modules and scripts) that
merely extend or make use of the Package, do not, by themselves, cause
the Package to be a Modified Version.  In addition, such works are not
considered parts of the Package itself, and are not subject to the
terms of this license.


General Provisions

(10)  Any use, modification, and distribution of the Standard or
Modified Versions is governed by this Artistic License. By using,
modifying or distributing the Package, you accept this license. Do not
use, modify, or distribute the Package, if you do not accept this
license.

(11)  If your Modified Version has been derived from a Modified
Version made by someone other than you, you are nevertheless required
to ensure that your Modified Version complies with the requirements of
this license.

(12)  This license does not grant you the right to use any trademark,
service mark, tradename, or logo of the Copyright Holder.

(13)  This license includes the non-exclusive, worldwide,
free-of-charge patent license to make, have made, use, offer to sell,
sell, import and otherwise transfer the Package with respect to any
patent claims licensable by the Copyright Holder that are necessarily
infringed by the Package. If you institute patent litigation
(including a cross-claim or counterclaim) against any party alleging
that the Package constitutes direct or contributory patent
infringement, then this Artistic License to you shall terminate on the
date that such litigation is filed.

(14)  Disclaimer of Warranty:
THE PACKAGE IS PROVIDED BY THE COPYRIGHT HOLDER AND CONTRIBUTORS "AS
IS' AND WITHOUT ANY EXPRESS OR IMPLIED WARRANTIES. THE IMPLIED
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, OR
NON-INFRINGEMENT ARE DISCLAIMED TO THE EXTENT PERMITTED BY YOUR LOCAL
LAW. UNLESS REQUIRED BY LAW, NO COPYRIGHT HOLDER OR CONTRIBUTOR WILL
BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, OR CONSEQUENTIAL
DAMAGES ARISING IN ANY WAY OUT OF THE USE OF THE PACKAGE, EVEN IF
ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

// This software is provided "AS IS" but the service it connects to has its own
// terms and polcies.
const notices = `
Use of the mktmpio service is subject to the the following:
 * mktmpio Privacy Policy
   https://mktmp.io/privacy-policy
 * mktmpio Terms of Service
   https://mktmp.io/terms-of-service
`

// The dependencies list below is manually populated using the output from:
// go list -f {{.Deps}} ./...
const thirdPartyShort = `
This binary includes 3rd party open source software.
`
const thirdpartyFull = `
This binary includes 3rd party software:

 * github.com/codegangsta/cli
   MIT license
   Copyright (C) 2013 Jeremy Saenz

 * github.com/docker/docker/pkg/stdcopy
   Apache License 2.0
   Copyright 2013-2015 Docker, Inc.

 * github.com/mitchellh/go-homedir
   MIT license
   Copyright (c) 2013 Mitchell Hashimoto

 * github.com/mktmpio/go-mktmpio
   Artistic-2.0 license
   Copyright (c) 2015 Datajin Technologies, Inc.

 * golang.org/x/crypto/ssh/terminal
   golang.org/x/net/websocket
   BSD-2 license
   Copyright (c) 2009 The Go Authors. All rights reserved.

 * gopkg.in/yaml.v2
   LGPL3 license
   Copyright (c) 2011-2014 - Canonical Inc.

 * github.com/Sirupsen/logrus
   MIT License
   Copyright (c) 2014 Simon Eskildsen
`

// This software is provided "AS IS"
const warrantiesShort = `
This software is provided "AS IS" and without warranty.
`
const warrantiesFull = `
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

// Contact info
const contact = `
For more information contact Datajin Technologies, Inc.:
 w: https://mktmp.io
 e: mktmpio@datajin.com
`

const summary = `
This is only a summary. For full warranty and license texts as well as the
complete listing of included 3rd party software, run with the --verbose flag.
`

// Definition for the 'mktmpio config' command
var LegalCommand = cli.Command{
	Name:   "legal",
	Usage:  "display licensing, copyright, and warranty notices",
	Action: licenseAction,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "verbose, v"},
	},
}

func licenseAction(c *cli.Context) {
	license := artisticShort
	thirdparty := thirdPartyShort
	warranties := warrantiesShort
	if c.Bool("verbose") {
		license = artisticFull
		thirdparty = thirdpartyFull
		warranties = warrantiesFull
	} else {
		c.App.Writer.Write([]byte(summary))
	}
	c.App.Writer.Write([]byte("\nLICENSE:\n" + license))
	c.App.Writer.Write([]byte("\nNOTICES:\n" + notices))
	c.App.Writer.Write([]byte("\nWARRANTIES:\n" + warranties))
	c.App.Writer.Write([]byte("\n3RD PARTY SOFTWARE:\n" + thirdparty))
	c.App.Writer.Write([]byte("\nCONTACT:\n" + contact))
}
