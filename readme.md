# Inspector

Inspector is a framework that allows build special code for inspecting values of arbitrary types with minimum or zero
allocations and using of reflect package.

It takes a path to the package as input argument and produces special struct with the suffix "Inspector". You may use
this struct for your issues.

No need to make it manually, there is a tool `inspc`. Just [use it](https://github.com/koykov/inspector/tree/master/inspc) to build inspectors.
