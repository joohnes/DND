<script lang="ts">
	export let showModal: Boolean;
	export let fullScreen: boolean = false;

	let dialog: HTMLDialogElement;

	$: if (dialog && showModal) dialog.showModal();
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
	bind:this={dialog}
	on:close={() => (showModal = false)}
	on:click|self={() => dialog.close()}
	class="modal"
>
<div class="modal-box grid justify-items-center {fullScreen ? "max-w-none w-full":""}">
    <slot />
	<form method="dialog" class="modal-backdrop">
		<button></button>
	</form>	
</dialog>
