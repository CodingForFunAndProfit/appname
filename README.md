# Go example project structure

I recreated the app structure mentioned in the article https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp. The Github for this project is https://github.com/benbjohnson/wtf and the series of articles start here https://www.gobeyond.dev/wtf-dial/. At this time there is only one introduction article. I liked his approach for structuring golang projects. I didn't like having all root domain objects (structs, interfaces, packages) in the rootfolder so I created an app folder that is the root of the program.

Additionally I wrote example code to gracefully shutdown the application running multiple go routines concurrently
