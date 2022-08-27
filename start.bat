echo "Start"
 rd /s /q docs
 mkdir docs
cp  -R build/html/. docs/
echo . > docs/.nojekyll