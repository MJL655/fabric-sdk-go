go mod vendor
cp vendor/modules.txt ../
cp -Rf vendor/github.com/studyzy/* ../../
cd vendor
rm -rf *
git checkout .
mv ../../modules.txt ./
rm -rf github.com/studyzy/*
mv -f ../../../crypto github.com/studyzy
cd ..
