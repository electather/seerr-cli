docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/open-api.yaml \
    -g go \
    -o /local/pkg/api \
    --additional-properties=packageName=api,isGoSubmodule=true,moduleName=seer-cli/pkg/api

# The generator may overwrite go.mod with a placeholder module name — fix it.
sed -i.bak 's|^module .*|module seer-cli/pkg/api|' pkg/api/go.mod && rm -f pkg/api/go.mod.bak