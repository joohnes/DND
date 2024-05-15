<script lang="ts">
	export let showModal: Boolean;

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
<div class="modal-box grid justify-items-center">
    <slot />
	<form method="dialog" class="modal-backdrop">
		<button></button>
	</form>	
</dialog>
