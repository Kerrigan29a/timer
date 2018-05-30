#!/bin/bash
# nothing to see here, just a utility i use to create new releases

CURRENT_VERSION=$(cat VERSION)
TO_UPDATE=(
    VERSION
)

echo -n "Current version is $CURRENT_VERSION, select new version: "
read NEW_VERSION
echo "Creating version $NEW_VERSION"

for file in "${TO_UPDATE[@]}"
do
    echo "    Patching $file"
    sed -i "s/$CURRENT_VERSION/$NEW_VERSION/g" $file
    git add $file
done
echo "DONE"

#git commit -m "Releasing $NEW_VERSION"
#git push

#git tag -a $NEW_VERSION -m "Release $NEW_VERSION"
#git push origin $NEW_VERSION
