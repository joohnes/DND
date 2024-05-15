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
		GetPlayerNames()
	});
	
	let rarityColorMap = new Map<string, string>([
		["Common", "badge-default"],
		["Rare", "badge-success"],
		["Epic", "badge-info"],
		["Legendary", "badge-warning"],
		["Artefact", "badge-error"]
	]);

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

	const TransferItemToPlayer = (e: any) => {
		const formData = new FormData(e.target)
		let playerName = formData.get("select-player")!.toString()

		let userId: string = "";
		for (const [id, name] of Object.entries(names)) {
			if (name === playerName) {
			userId = id;
			break;
			}
		}

		fetch(HOST + "bag/transfer/" + i.id + "/" + userId, {
			method: "POST",
		})
		window.location.reload()
	}

	const AddItemToPlayer = (e: any) => {
		const formData = new FormData(e.target)
		let playerName = formData.get("select-player")!.toString()

		let userId: string = "";
		for (const [id, name] of Object.entries(names)) {
			if (name === playerName) {
			userId = id;
			break;
			}
		}

		fetch(HOST + "player/" + userId + "/add-item/" + i.id, {
			method: "POST",
		})
		window.location.reload()
	}
	
	let names: {};
	async function GetPlayerNames() {
				const res = await fetch(HOST + 'players-names');
				const data = await res.json();
				names = data
	}

	let showModal = false;
	let modalBag = false;
	export let id: number = 0;
	export let bag: boolean = false;
	export let playerView: boolean = false;
</script>

{#if !modalBag}
<Modal bind:showModal>
	<span class="mb-4">Delete item?</span>
	<button class="btn btn-error btn-outline btn-md hover:btn-outline" on:click={()=> {DeleteItem()}}>DELETE</button>
</Modal>
{:else}
<Modal bind:showModal>
	{#if !bag}
		<button class="btn btn-outline btn-info btn-md mb-4 hover:btn-outline" on:click={()=>{TransferItemToBag()}}>To Bag</button>
	{/if}
	{#if bag || playerView}
		<button class="btn btn-info btn-md m-6" on:click={()=>{Drop()}}>Drop item</button>
	{/if}
	{#if names != null}
	<form on:submit|preventDefault={bag ? TransferItemToPlayer : AddItemToPlayer} class="flex flex-col justify-center">
	{#if bag}
		<select name="select-player" class="select select-bordered w-full max-w-xs mb-4">
			<option disabled selected>Select Player</option>
			{#each Object.entries(names) as id}
				<option>{id[1]}</option>
			{/each}
		</select>
	{:else}
		<select name="select-player" class="select select-bordered w-full max-w-xs mb-4">
			<option disabled selected>Select Player</option>
			{#each Object.entries(names) as id}
				<option>{id[1]}</option>
			{/each}
		</select>
	{/if}
	<button class="btn"><input type="submit" value="Change"></button>
	</form>
	
	{/if}
</Modal>
{/if}
<div>
	{#if i != undefined}
		<div class="item">
			<div class="info">
				<div>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-missing-attribute -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<span on:click={()=>{modalBag=true;showModal=true;}}>💰</span>
					<a href={window.location.origin+"/items/update/"+id}>⚙️</a>
				</div>
				<div>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-missing-attribute -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<span on:click={()=>{modalBag=false;showModal=true;}}>❌</span>
				</div>
			</div>
			<div class="cellHolder">
				<div class="underline underline-offset-1}">{i.name} x{i.quantity}</div>
				<div class="badge badge-ghost badge-outline {rarityColorMap.get(i.rarity)}">{i.rarity}</div>
			</div>
			<div class="cellHolder flex-col">
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
		@apply border-slate-500;
	}
	.cellHolder {
		display: flex;
		padding-top: 0.5rem;
		padding-bottom: 0.5rem;
		border-bottom: 1px solid;
		@apply border-slate-600;
		justify-content: space-between;
	}
	.statsCell {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}
	.info {
		display: flex;
		justify-content: space-between;
	}
	.info > div > span, .info > div > a {
		cursor: pointer;
	}
</style>
