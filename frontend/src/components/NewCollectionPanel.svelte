<script lang="ts">
  import { Panel, Page, Link, Navbar } from 'konsta/svelte'
	import type { CustomDropzoneProps, DropzoneEvent,RejectedFile } from 'svelte-dropzone-runes'
	import Dropzone from 'svelte-dropzone-runes'
  import { X, ImagePlus } from '@lucide/svelte'
  
  let { open, handleClose } = $props()
	let dropzoneElement: HTMLElement | undefined = $state()
	let isDraggingOver = $state(false)

	let files = $state({
		acceptedFiles: [] as File[],
		rejectedFiles: [] as RejectedFile<File>[]
	})

	function handleFilesSelect(e: DropzoneEvent<File>) {
		files = e
		isDraggingOver = false
	}
</script>

{#snippet CustomDropzone(props: CustomDropzoneProps)}
  <div
    bind:this={dropzoneElement}
    class={[
      `border-4 border-md-light-outline flex flex-col
      items-center justify-center gap-2 px-4 py-6 rounded-lg
      transition-colors duration-300`,
      isDraggingOver && "bg-brand-primary/10"
    ]}
    {...props}
  >
    <ImagePlus class="w-10 h-10 text-md-dark-on-surface opacity-50" />
    <p class="select-none text-md-dark-on-surface text-xs">Drop Here</p>
  </div>
{/snippet}

<Panel
  side="right"
  opened={open}
  onBackdropClick={handleClose}
  class="dark:bg-md-dark-surface"
>
  <Page class="bg-transparent">
    <Navbar
      title="New Collection"
      bgClass="bg-md-dark-surface"
      titleClass="font-medium select-none"
    >
      {#snippet right()}
        <Link
          iconOnly
          onClick={handleClose}
          class="hover:text-md-dark-on-surface transition-colors"
        >
          <X />
        </Link>
      {/snippet}
    </Navbar>
    <div class="p-4">
      <div>
        <p>Art</p>
        <Dropzone
          {dropzoneElement}
          onDrop={handleFilesSelect}
          onDragenter={() => {
            console.log('dragenter')
            isDraggingOver = true
          }}
          onDragleave={() => {
            isDraggingOver = false
          }}
          {CustomDropzone}
        />
      </div>
    </div>
  </Page>
</Panel>
