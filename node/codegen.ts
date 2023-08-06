import { CodegenConfig } from '@graphql-codegen/cli'
 
const config: CodegenConfig = {
  schema: 'api/*.gql',
  documents: [
    'api/fragment.gql'
  ],
  generates: {
    './src/graphql/': {
      preset: 'client'
    }
  }
}
 
export default config
