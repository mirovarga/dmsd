version=`git describe --tags --abbrev=0`
platforms=("darwin/arm64" "darwin/amd64" "linux/amd64")
other_files="LICENSE README.md"

for platform in "${platforms[@]}"; do
  platform_split=(${platform//\// })

  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  output_name=dmsd
  if [ $GOOS = "windows" ]; then
    output_name+='.exe'
  fi

  echo "Building for $platform"
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
	zip -qj9 dmsd-$version-$GOOS-$GOARCH.zip $output_name $other_files
	rm $output_name
done
