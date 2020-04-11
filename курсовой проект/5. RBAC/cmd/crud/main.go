package main

// Main CRUD package

import (
	_ "http"
	_ "rbac"
)

// Storing local packages in "vendor" folder, to allow import them with shorter paths
// https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go/35511866
