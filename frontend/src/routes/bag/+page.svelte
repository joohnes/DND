<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { writable } from 'svelte/store';
	import Item from '../../components/Item.svelte';
	import Modal from '../../components/Modal.svelte';
  	import { HOST } from "$lib/host";

    let holder: string;
	let names: {};
	let bagItems = writable([]);
	var menuItem: HTMLElement
	onMount(async () => {
		loadData()
	});

	async function loadData() {
		menuItem = document.getElementById("menu-bag")!
		if (menuItem != undefined) {
			menuItem.classList.remove("btn-outline")
		}
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
		GetPlayerNames()
	}

	onDestroy(() => {
		if (menuItem != undefined) {
			menuItem.classList.add("btn-outline")
		}
	})

	const ChangeHolder = (e: any) => {
		const formData = new FormData(e.target)
		let playerName = formData.get("select-player")!.toString()

		let userId: string = "";
		for (const [id, name] of Object.entries(names)) {
			if (name === playerName) {
			userId = id;
			break;
			}
		}

		fetch(HOST + "bag/holder", {
			method: "POST",
			body: JSON.stringify(parseInt(userId))
		})
		window.location.reload()
	}

	async function GetPlayerNames() {
				const res = await fetch(HOST + 'players-names');
				const data = await res.json();
				names = data
	}

	let restartKey = {};
	const restart = () => {
		restartKey = {}
		loadData()
	}

	let showModal = false;
</script>

<Modal bind:showModal>
	{#if names != null}
	<form on:submit|preventDefault={ChangeHolder} class="flex flex-col justify-center">
		<select name="select-player" class="select select-bordered w-full max-w-xs mb-4">
			<option disabled selected>Select Player</option>
			{#each Object.values(names) as name}
			<option>{name}</option>
			{/each}
		</select>
		<button class="btn"><input type="submit" value="Change"></button>
	</form>
	{:else}
	No Players!
	{/if}
</Modal>
<div class="flex justify-center pb-6">
	<span class="text-3xl my-auto">Bag of Holding</span>
	<button class="btn btn-outline btn-success ml-4 rounded-full" on:click={()=>{showModal=true;}}>Change holder</button>
</div>
{#if holder != undefined && holder != ""}
<div class="flex justify-center pb-6">
	<span class="text-3xl">Holder: {holder}</span>
</div>
{/if}

<div class="flex flex-wrap gap-5 justify-center">
	{#if $bagItems != null}
	{#each $bagItems as item}
		<Item id={item} bag={true} restart={restart}/>
	{/each}
	{:else}
	<h1>No Items</h1>
	{/if}
</div>
