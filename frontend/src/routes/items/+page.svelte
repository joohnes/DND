<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { writable } from 'svelte/store';
	import Item from '../../components/Item.svelte';
	import { HOST } from '$lib/host';

	let items = writable([]);
	var menuItem: HTMLElement
	onMount(async () => {
		menuItem = document.getElementById("menu-items")!
		if (menuItem != undefined) {
			menuItem.classList.remove("btn-ghost")
			menuItem.classList.add("btn-remove")
			menuItem.classList.add("btn-info")
		}

		const res = await fetch(HOST + 'items');
		const data = await res.json();
		items.set(data);
	});

	onDestroy(() => {
		if (menuItem != undefined) {
			menuItem.classList.add("btn-ghost")
			menuItem.classList.add("btn-outline")
			menuItem.classList.remove("btn-info")
		}
	})
</script>

{#if $items != null}
	<div class="flex justify-center pb-6">
		<span class="text-3xl">Items</span>
	</div>
	<div class="flex flex-wrap gap-5 justify-center">
		{#each $items as item}
			<Item id={item} />
		{/each}
	</div>
	{:else}
		<div class="flex justify-center pb-6">
			<span class="text-3xl">No Items</span>
		</div>
{/if}
