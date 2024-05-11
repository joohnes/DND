<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	import Item from '../../components/Item.svelte';
	import Modal from '../../components/Modal.svelte';
  	import { HOST } from "$lib/host";

    let holder: string;
	let bagItems = writable([]);
	onMount(async () => {
		{
			const res = await fetch(HOST + "bag");
			const data = await res.json();
			bagItems.set(data.items);
		}
		{
			const res = await fetch(HOST + "bag/holder")
			const data = await res.json()
			holder = data
		}
	});

	async function ChangeHolder(id: string) {
		await fetch(HOST + "bag/holder", {
			method: "POST",
			body: JSON.stringify(id)
		})
	}

	let names: [string, unknown][] | null;
	async function GetPlayerNames() {
				const res = await fetch(HOST + 'players-names');
				const data = await res.json();
				names = Object.entries(data)
	}
	GetPlayerNames()

	let showModal = false;
</script>

<Modal bind:showModal>
	<h2>New Holder:</h2>
	{#if names != null}
	{#each names as id}
		<button on:click={()=>{ChangeHolder(id[0])}}>{id[1]}</button>
	{/each}
	{/if}
</Modal>
<h1>Bag of Holding</h1>
{#if holder != undefined}
<h2>Current Holder: {holder}</h2>
{/if}
<button on:click={()=>{showModal=true;}}>Change owner</button>
<div class="ItemHolder">
	{#if $bagItems != null}
	{#each $bagItems as item}
		<Item id={item} bag={true} />
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
</style>