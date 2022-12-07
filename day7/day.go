package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

type node struct {
	parent   *node
	subnodes []*node
	name     string
	size     int
	nodetype int // 0 - dir, 1 - file

}

// Check if any character appears more than once in s.s
func (n *node) addNode(newnode *node) {
	n.subnodes = append(n.subnodes, newnode)
}

func (n *node) getSize() int {
	if n.nodetype == 1 {
		return n.size
	}
	size := 0
	for _, subnode := range n.subnodes {
		size += subnode.getSize()
	}
	return size
}

func (n *node) getSizeAnswer1(r *int) int {
	if n.nodetype == 1 {
		return n.size
	}
	size := 0
	for _, subnode := range n.subnodes {
		size += subnode.getSizeAnswer1(r)
	}
	if size <= 100000 {
		*r += size
		// println(n.name, size)
	}
	return size
}

func (n *node) getSizeAnswer2(min *int, threshold int) int {
	if n.nodetype == 1 {
		return n.size
	}
	size := 0
	for _, subnode := range n.subnodes {
		size += subnode.getSizeAnswer2(min, threshold)
	}
	if size >= threshold {
		// *r += size
		if size < *min {
			*min = size
		}
		// println(n.name, size)
	}
	return size
}

// func (n *node) getSizeAnswer2(d []node) int {
// 	if n.nodetype == 1 {
// 		return n.size
// 	}
// 	size := 0
// 	for _, subnode := range n.subnodes {
// 		size += subnode.getSizeAnswer2(d)
// 	}
// 	d = append(d, *n)
// 	println(n.name, size)
// 	return size
// }

func main() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// println(len(b))
	// c := tj.DoRegex("123-45", `(\d+)-(\d+)`)
	// for _, i := range c {
	// 	println(i)
	// }
	// return
	// for r := range b[0] {
	// 	println(string(rune(r)))
	// }
	// input := b[0]

	pwd := &node{name: "/", nodetype: 0}
	root := pwd
	for _, line := range b {
		split := strings.Split(line, " ")
		if split[0] == "$" {
			if split[1] == "cd" {
				// cd command
				if split[2] == ".." {
					pwd = pwd.parent
				} else {
					for _, subnode := range pwd.subnodes {
						if subnode.name == split[2] {
							pwd = subnode
							break
						}
					}
				}
			} else if split[1] == "ls" {
				// ls command
				// println(split)
			}
		} else if split[0] == "dir" {
			// list directory
			pwd.addNode(&node{name: split[1], nodetype: 0, parent: pwd})
		} else {
			// list file
			newNode := &node{name: split[1], nodetype: 1, parent: pwd, size: tj.Str2int(split[0])}
			pwd.addNode(newNode)
		}

		// println(line)
		// println(pwd.name)
	}
	total := 0
	root.getSizeAnswer1(&total)
	println("Part 1:", total)

	println("/ size:")
	println(root.getSize())
	println("Empty space:")
	empty := 70000000 - root.getSize()
	println(empty)
	freeup := 30000000 - empty
	println("We need to free up", freeup)

	min := 70000000
	root.getSizeAnswer2(&min, freeup)
	println("Min size:", min)

}
