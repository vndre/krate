<script lang="ts">
  import { Button } from 'konsta/svelte'

  let { children, ...props } = $props()
</script>

<Button
  {...props}
  class="font-medium max-w-max"
  rounded
>
  {@render children()}
</Button>
