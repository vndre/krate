import { error } from '@sveltejs/kit'
import type { PageLoad } from './$types'
import { GetCollections } from '@wailsjs/handlers/CollectionHandler'
import type { collection_model } from '@wailsjs/models'

export const load: PageLoad = async () => {
  try {
    // const collections = await GetCollections()
    // return {
    //   collections,
    // }
    return {
      collections: [],
    }
  } catch (e) {
    return error(500, e instanceof Error ? e.message : 'Failed to load collections')
  }
}
