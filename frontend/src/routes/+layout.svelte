<script lang="ts">
  import { App, KonstaProvider } from 'konsta/svelte'
  
	import '../app.css'
	import favicon from '$lib/assets/favicon.svg'

	let { children } = $props()
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<KonstaProvider theme="material">
  <App theme="material" class="k-material p-16 flex">
    {@render children()}
  </App>
</KonstaProvider>