package main

import (
	"github.com/vocdoni/arbo"
	"go.vocdoni.io/dvote/db"
	badb "go.vocdoni.io/dvote/db/badgerdb"
)

func (app *App) createPoseidonTreeDb(dbname string, levels int, sha256 bool) (*badb.BadgerDB, *MerkleTree, error) {
	// Create a new database
	var opts db.Options
	opts.Path = dbname
	dbpoint, err := badb.New(opts)
	if err != nil {
		logs.logError("Failed to access database "+dbname+" !!!", err)
		return nil, nil, err
	}
	return app.createPoseidonTree(dbpoint)
}

func (app *App) createTreeDb(dbname string, levels int, sha256 bool) (*badb.BadgerDB, *arbo.Tree, error) {
	// Create a new database
	var opts db.Options
	opts.Path = dbname
	dbpoint, err := badb.New(opts)
	if err != nil {
		logs.logError("Failed to access database "+dbname+" !!!", err)
		return nil, nil, err
	}
	return app.createTree(dbpoint, levels, sha256)
}

// createTree creates a new Merkle tree associated with the given database.
func (app *App) createPoseidonTree(dbpoint *badb.BadgerDB) (*badb.BadgerDB, *MerkleTree, error) {
	// Create a new Merkle tree instance
	tree := NewMerkleTree(dbpoint)
	/*
		if tree == nil {
			//err := "failed to create Merkle tree"
			logs.logError("Failed to create tree !!!", nil)
			return dbpoint, nil, err
		}
	*/
	return dbpoint, tree, nil
}

func (app *App) createTree(dbpoint *badb.BadgerDB, levels int, sha256 bool) (*badb.BadgerDB, *arbo.Tree, error) {
	// Create a new tree associated with the database
	var config arbo.Config

	config = arbo.Config{
		Database:     dbpoint,
		MaxLevels:    levels,
		HashFunction: arbo.HashFunctionBlake2b}

	if sha256 {
		config.HashFunction = arbo.HashFunctionSha256
	}

	Tree, err := arbo.NewTree(config)

	if err != nil {
		logs.logError("Failed to create tree !!!", err)
		return dbpoint, nil, err
	}

	return dbpoint, Tree, nil
}
