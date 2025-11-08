<script lang="ts">
	import { page } from '$app/state'
</script>

<div class="flex flex-col items-center justify-center h-full m-auto gap-4 select-none">
  <h1 class="font-bold text-8xl">{page.status}</h1>
  <h2>{page?.error?.message}</h2>
</div>