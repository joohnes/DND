<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	import Item from '../../components/Item.svelte';
	import { HOST } from '$lib/host';

	let items = writable([]);
	onMount(async () => {
		const res = await fetch(HOST + 'items');
		const data = await res.json();
		items.set(data);
	});
</script>

<div style="display: flex;">
	<h1>Items</h1>
	<button class="btn" on:click={()=>{window.location.href = window.location.origin+"/items/create"}}>Create Item</button>
</div>
<div class="ItemHolder">
	{#if $items != null}
	{#each $items as item}
		<Item id={item} />
	{/each}
	{:else}
	<h1>No Items</h1>
	{/if}
</div>

<style>
	.ItemHolder {
		display: flex;
		flex-wrap: wrap;
		gap: 10px;
	}

	.btn {
		margin: auto;
		margin-left: 1rem;
	}
</style>
