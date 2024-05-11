<script lang="ts">
	import { onMount } from 'svelte';
	import Modal from "./Modal.svelte"
  	import { HOST } from "$lib/host";

	let i: Item;
	onMount(async function () {
		const res = await fetch(HOST + 'item/' + id);
		const data = await res.json();
		i = {
			id: data.id,
			name: data.name,
			description: data.description,
			ability: data.ability,
			rarity: data.rarity,
			strength: data.strength,
			endurance: data.endurance,
			perception: data.perception,
			intelligence: data.intelligence,
			agility: data.agility,
			accuracy: data.accuracy,
			charisma: data.charisma,
			quantity: data.quantity
		};
	});

	async function Drop() {
		if (!playerView) {
			await fetch(HOST + "bag/drop/" + i.id, {
				method: "DELETE",
			})
			window.location.reload()
		} else {
			await fetch(HOST + "player/drop-item/" + i.id, {
				method: "DELETE",
			})
			window.location.href = window.location.origin + "/players"
		}
	}
	async function DeleteItem() {
		await fetch(HOST + "item/" + i.id, {
        method: "DELETE",
		})
		window.location.reload()
	}

	async function TransferItemToBag() {
		await fetch(HOST + "bag/add/" + i.id, {
			method: "POST",
		})
		window.location.reload()
	}

	async function TransferItemToPlayer(id: string) {
		await fetch(HOST + "bag/transfer/" + i.id + "/" + id, {
			method: "POST",
		})
		window.location.reload()
	}

	async function AddItemToPlayer(id: string) {
		await fetch(HOST + "player/" + id + "/add-item/" + i.id, {
			method: "POST",
		})
		window.location.reload()
	}
	
	let names: [string, unknown][] | null;
	async function GetPlayerNames() {
				const res = await fetch(HOST + 'players-names');
				const data = await res.json();
				names = Object.entries(data)
	}
	GetPlayerNames()

	let showModal = false;
	let modalBag = false;
	export let id: number = 0;
	export let bag: boolean = false;
	export let playerView: boolean = false;
</script>
{#if !modalBag}
<Modal bind:showModal>
	<h2>Delete item?</h2>
	<button on:click={()=> {DeleteItem()}}>DELETE</button>
</Modal>
{:else}
<Modal bind:showModal>
	<h2>Transfer item</h2>
	{#if !bag}
		<button on:click={()=>{TransferItemToBag()}}>To Bag</button>
	{/if}
	{#if bag || playerView}
		<h3><button on:click={()=>{Drop()}}>Drop</button></h3>
	{/if}
	{#if names != null}
		<h3>To Player</h3>
	{#if bag}
		{#each names as id}
			<button on:click={()=>{TransferItemToPlayer(id[0])}}>{id[1]}</button>
		{/each}
	{:else}
		{#each names as id}
			<button on:click={()=>{AddItemToPlayer(id[0])}}>{id[1]}</button>
		{/each}
	{/if}
	
	{/if}
</Modal>
{/if}
<div>
	{#if i != undefined}
		<div class="item">
			<div class="info">
				<div class="type">Item</div>
				<div>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-missing-attribute -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<a on:click={()=>{modalBag=true;showModal=true;}}>💰</a>
					<a href={window.location.origin+"/items/update/"+id}>⚙️</a>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-missing-attribute -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<a on:click={()=>{modalBag=false;showModal=true;}}>❌</a>
				</div>
			</div>
			<div class="cellHolder">
				<div>{i.name} x{i.quantity}</div>
				<div>{i.rarity}</div>
			</div>
			<div class="cellHolder stats">
				<div class="statsCell">
					<div>Strength:</div>
					<div>{i.strength}</div>
				</div>
				<div class="statsCell">
					<div>Endurance:</div>
					<div>{i.endurance}</div>
				</div>
				<div class="statsCell">
					<div>Perception:</div>
					<div>{i.perception}</div>
				</div>
				<div class="statsCell">
					<div>Intelligence:</div>
					<div>{i.intelligence}</div>
				</div>
				<div class="statsCell">
					<div>Agility:</div>
					<div>{i.agility}</div>
				</div>
				<div class="statsCell">
					<div>Accuracy:</div>
					<div>{i.accuracy}</div>
				</div>
				<div class="statsCell">
					<div>Charisma:</div>
					<div>{i.charisma}</div>
				</div>
			</div>
			<div class="cellHolder">
				<div>{i.description}</div>
			</div>
			<div class="cellHolder">
				<div>{i.ability}</div>
			</div>
		</div>
	{/if}
</div>

<style>
	a {
		text-decoration: none;
		color: #2d2d2d;
		font-weight: 400;
	}
	.item {
		display: flex;
		flex-direction: column;
		width: 15rem;
		justify-content: space-between;
		padding: 1rem;
		border: 3px solid black;
		border-radius: 1rem;
	}
	.cellHolder {
		display: flex;
		border-bottom: 1px solid rgba(0, 0, 0, 0.311);
		justify-content: space-between;
	}
	.stats {
		flex-direction: column;
	}
	.statsCell {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}
	.type {
		margin-top: auto;
		font-size: 0.6rem;
	}
	.info {
		display: flex;
		justify-content: space-between;
	}
</style>
