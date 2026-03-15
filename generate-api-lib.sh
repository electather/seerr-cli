docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/open-api.yaml \
    -g go \
    -o /local/pkg/api \
    --additional-properties=packageName=api,isGoSubmodule=true,moduleName=seerr-cli/pkg/api

# The generator may overwrite go.mod with a placeholder module name — fix it.
sed -i.bak 's|^module .*|module seerr-cli/pkg/api|' pkg/api/go.mod && rm -f pkg/api/go.mod.bak