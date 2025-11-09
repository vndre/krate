<script lang="ts">
  import { Panel, Page } from 'konsta/svelte'
	import type { CustomDropzoneProps, DropzoneEvent,RejectedFile } from 'svelte-dropzone-runes'
	import Dropzone from 'svelte-dropzone-runes'
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
  <div bind:this={dropzoneElement} {...props}>Custom Dropzone</div>
{/snippet}

<Panel
  side="left"
  opened={open}
  onBackdropClick={handleClose}
>
  <Page>
    <Dropzone {dropzoneElement} onDrop={handleFilesSelect} {CustomDropzone} />
  </Page>
</Panel>
