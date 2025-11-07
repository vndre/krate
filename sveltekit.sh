manager=$1


sed -i "s|npm|$manager|g" wails.json
sed -i 's|"auto",|"auto",\n  "wailsjsdir": "./frontend/src/lib",|' wails.json
sed -i "s|all:frontend/dist|all:frontend/build|" main.go
rm -r frontend
$manager create svelte@latest frontend
cd frontend
$manager i
$manager uninstall @sveltejs/adapter-auto
$manager i -D @sveltejs/adapter-static
echo -e "export const prerender = true\nexport const ssr = false" > src/routes/+layout.ts
sed -i "s|-auto';|-static';|" svelte.config.js
cd ..
wails dev