/*
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// this file was generated by tool/embed.sh: don't modify!

package main

import "strings"

func help(arg string) string {
	switch strings.ToLower(arg) {
		case "help":
			return `
$1Usage$0
-----

    souschef command [--flags]

$1Commands$0
--------

    $1init$0     initialise new Sous Chef project
    $1order$0    create a new order
    $1list$0     show the list of current orders
    $1render$0   start the render queue
    $1clean$0    remove finished orders
    $1help$0     print this message and others

Use $1souschef help [command]$0 for more information on each of 
the above.
`
	}
	return help("help")
}