module example.com/solo/hellogo

go 1.22.2

require example.com/The-solo/mystrings v0.0.0
replace example.com/The-solo/mystrings v0.0.0 => ../mystrings

// 'Replace' keyword only works for the local packages 
