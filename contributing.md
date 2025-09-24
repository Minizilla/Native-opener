Si c'est un repo dérivés (fork, clone) Sur votre github https://github.com/settings/tokens/new?scopes=repo,write:packages pour créer 'mytoken'

Dans votre config terminal :
export GITHUB_TOKEN=mytoken
Redémarrer le terminal

Codez...

git tag -a v0.1.0 -m "My release"
git push origin v0.1.0

goreleaser release --snapshot --clean
