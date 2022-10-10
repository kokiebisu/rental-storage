package main

type Event struct {
	Field string `json:"field"`
	Arguments interface {} `json:"arguments"`
	Identity Identity `json:"identity"`
}

type Identity struct {
	ResolverContext ResolverContext
}

type ResolverContext struct {
	UId string `json:"uid"`
}