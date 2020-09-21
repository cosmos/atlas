package server_test

// func TestController(t *testing.T) {
// 	dsn := "user=aleksbez dbname=atlas_local port=5432 sslmode=disable"

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	require.NoError(t, err)

// 	// var (
// 	// 	mod server.Module
// 	// )

// 	// db.Find(&mod, "name = 'x/foo'")

// 	// a := db.Model(&mod).Association("Bug")
// 	// require.NoError(t, a.Find(&mod.Bug))

// 	// fmt.Println(mod)

// 	bugTracker := server.BugTracker{URL: "cat", Contact: "cat"}
// 	keywords := []server.Keyword{
// 		{Name: "cat"},
// 		{Name: "bar"},
// 	}

// 	mod1 := &server.Module{
// 		Name:        "x/cat5",
// 		Description: "Provides a module that exposes token transfer functionality.",
// 		Homepage:    "https://github.com/cosmos/cosmos-sdk/tree/master/x/cat",
// 		Repo:        "https://github.com/cosmos/cosmos-sdk",
// 		BugTracker:  bugTracker,
// 		Keywords:    keywords,
// 		Versions: []server.ModuleVersion{
// 			{Version: "v1.0.0"},
// 		},
// 	}

// 	mod2 := &server.Module{
// 		Name:        "x/cat6",
// 		Description: "Provides a module that exposes token transfer functionality.",
// 		Homepage:    "https://github.com/cosmos/cosmos-sdk/tree/master/x/cat",
// 		Repo:        "https://github.com/cosmos/cosmos-sdk",
// 		BugTracker:  bugTracker,
// 		Keywords:    keywords,
// 		Versions: []server.ModuleVersion{
// 			{Version: "v2.0.0"},
// 		},
// 	}

// 	user := &server.User{
// 		Name: "author1", Email: "author1@email.com",
// 		Modules: []server.Module{
// 			*mod1,
// 			*mod2,
// 		},
// 	}
// 	fmt.Println(user)
// 	fmt.Println("============================")

// 	var mod server.Module
// 	db.Find(&mod, "name = 'x/cat'")
// 	db.Model(&mod).Association("Author").Find(&mod.Author)
// 	fmt.Println("AUTHOR:", mod.Author)

// 	// require.NoError(t, db.Model(&server.User{}).Where("name = ?", "author1").Debug().Update("modules", user.Modules).Error)
// 	// require.NoError(t, db.Debug().Save(user).Error)
// 	// require.NoError(t, db.Debug().Save(mod2).Error)

// 	// require.NoError(t, db.Debug().AutoMigrate(
// 	// 	&server.ModuleKeywords{},
// 	// 	&server.User{},
// 	// 	&server.Module{},
// 	// 	&server.BugTracker{},
// 	// 	&server.ModuleVersion{},
// 	// 	&server.Keyword{},
// 	// ))
// }
