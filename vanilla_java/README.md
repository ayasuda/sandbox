Hello, World sample on vanilla JAVA.

## How to build

First, you should download library.

```console
$ wget https://repo1.maven.org/maven2/org/yaml/snakeyaml/1.21/snakeyaml-1.21.jar
```

Second, build and run it.

```console
$ javac -classpath .:snakeyaml-1.21.jar Application.java # compile
$ java -classpath .:snakeyaml-1.21.jar Application # run
$ jar cvfm app.jar Manifest.txt Application.class generators/*.class # packaging
$ java -jar app.jar # run via package
```
