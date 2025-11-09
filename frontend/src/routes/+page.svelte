<script lang="ts">
  import Button from '@components/Button.svelte'
  import NewCollectionPanel from '@components/NewCollectionPanel.svelte'
	import type { PageProps } from './$types'
  
	let { data }: PageProps = $props()
  const { collections } = data

  let newCollectionPanelOpen = $state(false)

  function handleToggleNewCollectionPanel() {
    newCollectionPanelOpen = !newCollectionPanelOpen
  }
  
  // async function handleCreate() {
  //   // Example: Create a new collection
  //   try {
  //     await CreateCollection('New Collection', 'Action', '', 0)
  //     await loadCollections() // Reload the list
  //   } catch (e) {
  //     error = e instanceof Error ? e.message : 'Failed to create collection'
  //     console.error('Error creating collection:', e)
  //   }
  // }
</script>

<div class="flex flex-col gap-4 w-full">
  <div class="flex items-center justify-between">
    <h1 class="text-4xl font-bold select-none">Collections</h1>
    <Button onclick={handleToggleNewCollectionPanel}>+ New</Button>
  </div>

  {#if collections.length === 0}
    <div class="text-neutral-500 m-auto text-2xl select-none">No collections found. Create one to get started.</div>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {#each collections as collection (collection.Id)}
        <div class="border rounded-lg p-4 hover:shadow-lg transition-shadow">
          {#if collection.ImageUrl}
            <img src={collection.ImageUrl} alt={collection.Name} class="w-full h-48 object-cover rounded mb-2" />
          {/if}
          <h2 class="text-xl font-semibold">{collection.Name}</h2>
          <p class="text-gray-600 text-sm">{collection.Genre}</p>
          <div class="mt-2">
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="bg-blue-600 h-2 rounded-full" 
                style="width: {collection.Progress}%"
              ></div>
            </div>
            <p class="text-xs text-gray-500 mt-1">{collection.Progress}% complete</p>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<NewCollectionPanel open={!newCollectionPanelOpen} handleClose={handleToggleNewCollectionPanel} />