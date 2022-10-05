package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type User struct {
	name string
	ID   int
}

func (u *User) AddUser(name string, ID int) *User {
	u.name = name
	u.ID = ID
	return u
}

func printUser(u User) {
	fmt.Println("Name:", u.name, "\tID:", u.ID)
}

type Transaction struct {
	From   User
	To     User
	Amount int
}

func (t *Transaction) AddTransaction(from User, to User, amount int) *Transaction {
	t.From = from
	t.To = to
	t.Amount = amount
	return t
}

func printTransaction(t Transaction) {
	fmt.Println("From:", t.From.name, "\tTo:", t.To.name, "\tAmount:", t.Amount)
}

type Node struct {
	Index        int
	Transactions []Transaction
	Hash         string
	PrevHash     string
	Nonce        string
}

func (n *Node) AddNode(index int, transactions []Transaction, prevHash string, nonce string) *Node {
	n.Index = index
	n.Transactions = transactions
	n.PrevHash = prevHash
	n.Nonce = nonce
	n.Hash = n.calculateHash()
	return n
}

func printNode(n Node) {
	fmt.Println("Index:", n.Index, "\tPrevHash:", n.PrevHash, "\tHash:", n.Hash, "\tNonce:", n.Nonce)
	for _, v := range n.Transactions {
		printTransaction(v)
	}
}

type merkleNode struct {
	Left  *merkleNode
	Right *merkleNode
	Data  Node
}

func (m *merkleNode) AddMerkleNode(left *merkleNode, right *merkleNode, data Node) *merkleNode {
	m.Left = left
	m.Right = right
	m.Data = data
	return m
}

func (m *merkleNode) AddNode(node Node) *merkleNode {
	if m.Left == nil && m.Right == nil {
		return new(merkleNode).AddMerkleNode(nil, nil, node)
	}
	if m.Left != nil {
		m.Left = m.Left.AddNode(node)
	} else {
		m.Right = m.Right.AddNode(node)
	}
	return m
}
func printMerkeNode(m *merkleNode) {
	if m == nil {
		return
	}
	printNode(m.Data)
	printMerkeNode(m.Left)
	printMerkeNode(m.Right)
}

type Blockchain struct {
	root *merkleNode
}

func printBlockchain(b Blockchain) {
	printMerkeNode(b.root)
}

// add node to blockchain
func (b *Blockchain) AddNode(node Node) {
	//add node to merkle tree
	if b.root == nil {
		b.root = new(merkleNode).AddMerkleNode(nil, nil, node)
	} else {
		b.root = b.root.AddNode(node)
	}
}

func (n *Node) calculateHash() string {
	//calculate hash of node
	string := strconv.Itoa(n.Index) + n.PrevHash + n.Nonce
	for _, v := range n.Transactions {
		string += v.From.name + v.To.name + strconv.Itoa(v.Amount)
	}
	hash := sha256.Sum256([]byte(string))
	return fmt.Sprintf("%x", hash)
}

func main() {
	//create blockchain
	b := Blockchain{nil}
	//create users
	u1 := new(User).AddUser("Bob", 1)
	u2 := new(User).AddUser("Alice", 2)
	u3 := new(User).AddUser("John", 3)
	//create transactions
	t1 := new(Transaction).AddTransaction(*u1, *u2, 10)
	t2 := new(Transaction).AddTransaction(*u2, *u3, 20)
	t3 := new(Transaction).AddTransaction(*u3, *u1, 30)
	//create nodes
	n1 := new(Node).AddNode(1, []Transaction{*t1, *t2}, "", "")
	n2 := new(Node).AddNode(2, []Transaction{*t3}, n1.Hash, "")
	//add nodes to blockchain
	b.AddNode(*n1)
	b.AddNode(*n2)

	//print blockchain
}
