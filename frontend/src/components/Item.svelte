<script lang="ts">
	import { onMount } from 'svelte';
	import Modal from "./Modal.svelte"
  	import { HOST } from "$lib/host";

	onMount(async function () {
		GetPlayerNames()
	});
	
	let rarityColorMap = new Map<string, string>([
		["Common", "badge-default"],
		["Rare", "badge-success"],
		["Epic", "badge-info"],
		["Legendary", "badge-warning"],
		["Artefact", "badge-error"]
	]);

	let slotMap = new Map<number, string>([
		[1, "Head"],
		[2, "Chest"],
		[3, "Shoulders"],
		[4, "Hands"],
		[5, "Legs"],
		[6, "Feet"],
	])

	async function Drop() {
		if (!playerView) {
			await fetch(HOST + "bag/drop/" + item.id, {
				method: "DELETE",
			})
			showModal=false;
			restart()
		} else {
			await fetch(HOST + "player/drop-item/" + item.id, {
				method: "DELETE",
			})
			window.location.href = window.location.origin + "/players"
		}
	}
	async function DeleteItem() {
		await fetch(HOST + "item/" + item.id, {
        method: "DELETE",
		})
		showModal=false;
		restart()
	}

	async function TransferItemToBag() {
		await fetch(HOST + "bag/add/" + item.id, {
			method: "POST",
		})
		showModal=false;
		restart()
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

		fetch(HOST + "bag/transfer/" + item.id + "/" + userId, {
			method: "POST",
		})
		showModal=false;
		restart()
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

		fetch(HOST + "player/" + userId + "/add-item/" + item.id, {
			method: "POST",
		})
		showModal=false;
		restart()
	}
	
	let names: {};
	async function GetPlayerNames() {
				const res = await fetch(HOST + 'players-names');
				const data = await res.json();
				names = data
	}

	async function Equip() {
		const data = {
			playerID: playerID,
			itemID: item.id,
		}

		await fetch(HOST + "item/equip", {
			method: "POST",
			body: JSON.stringify(data)
		})
		restart()
	}

	async function Unequip() {
		await fetch(HOST + "item/unequip/" + item.id, {
			method: "DELETE",
		})
		restart()
	}

	let showModal = false;
	let modalBag = false;
	export let item: Item;
	export let playerID: number = 0;
	export let bag: boolean = false;
	export let playerView: boolean = false;
	export let eqView: boolean = false;
	export let equipped: boolean = false;
	export let restart: Function;
</script>

{#key showModal}
	{#if !modalBag}
	<Modal bind:showModal>
		<span class="mb-4">Delete item?</span>
		<button class="btn btn-error btn-outline btn-md hover:btn-outline" on:click={()=> {DeleteItem()}}>DELETE</button>
	</Modal>
	{:else}
	<Modal bind:showModal>
		<div class="flex gap-24">
			<div class="flex flex-col justify-start content-start">
			{#if !bag}
					<button class="btn btn-outline btn-info btn-md mb-4 hover:btn-outline px-8" on:click={()=>{TransferItemToBag()}}>To Bag</button>
				{/if}
			{#if bag || playerView || eqView}
				{#if eqView}
					<div class="flex flex-col justify-center">
						{#if !equipped}
						<button on:click={Equip} class="btn btn-outline btn-accent btn-md hover:btn-outline mb-4 px-8">Equip</button>
						{:else}
						<button on:click={Unequip} class="btn btn-outline btn-error btn-md hover:btn-outline mb-4 px-8">Unequip</button>
						{/if}
					</div>
				{/if}
				{#if bag || playerView}
					<button class="btn btn-outline btn-error btn-md px-8" on:click={()=>{Drop()}}>Drop item</button>
				{/if}
			{/if}
		</div>
			{#if names != null}
			<div class="flex flex-col justify-start content-start">
				<form on:submit|preventDefault={bag ? TransferItemToPlayer : AddItemToPlayer} class="flex flex-col justify-center">
					<select name="select-player" class="select select-bordered w-full max-w-xs mb-4">
						<option disabled selected>Select Player</option>
						{#each Object.entries(names) as id}
							<option>{id[1]}</option>
						{/each}
					</select>
					<input class="btn" type="submit" value="Change">
				</form>
			</div>
			{/if}
		</div>
	</Modal>
	{/if}
	{#if item != undefined}
	<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
		<div class="item collapse cursor" tabindex="0">
			<input id="open" type="checkbox"/>
			<div class="flex justify-between collapse-title" style="padding-inline-end: 1rem!important;">
				<div class="grow">
					<div class="underline underline-offset-4">{item.name} x{item.quantity}</div>
					{#if item.owner != ""}
						<div class="name">Owner: {item.owner}</div>
					{/if}
				</div>
				<div style="z-index: 999!important;">
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-missing-attribute -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<span class="cursor" on:click={()=>{modalBag=true;showModal=true;}}>💰</span>
					<a class="cursor" href={window.location.origin+"/items/update/"+item.id}>⚙️</a>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-missing-attribute -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<span class="cursor" on:click={()=>{modalBag=false;showModal=true;}}>❌</span>
				</div>
			</div>
			<div class="collapse-content">
				<div class="cellHolder">
					<div class="badge badge-ghost badge-outline {rarityColorMap.get(item.rarity)}">{item.rarity}</div>
					{#if item.slot != undefined}
						<div class="badge badge-ghost badge-outline">{slotMap.get(item.slot)}</div>
					{/if}
				</div>
				<div class="cellHolder">
					<div class="badge badge-outline badge-accent my-auto">{item.permille}‰</div>
					<div>
						<div class="badge badge-secondary badge-outline">ATK {item.attack}</div>
						<div class="badge badge-info badge-outline">DEF {item.defense}</div>
					</div>
				</div>
				<div class="cellHolder flex-col">
					{#if item.strength > 0}
					<div class="statsCell">
						<div>Strength:</div>
						<div>{item.strength}</div>
					</div>
					{/if}
					{#if item.endurance > 0}
					<div class="statsCell">
						<div>Endurance:</div>
						<div>{item.endurance}</div>
					</div>
					{/if}
					{#if item.perception > 0}
					<div class="statsCell">
						<div>Perception:</div>
						<div>{item.perception}</div>
					</div>
					{/if}
					{#if item.intelligence > 0}
					<div class="statsCell">
						<div>Intelligence:</div>
						<div>{item.intelligence}</div>
					</div>
					{/if}
					{#if item.agility > 0}
					<div class="statsCell">
						<div>Agility:</div>
						<div>{item.agility}</div>
					</div>
					{/if}
					{#if item.accuracy > 0}
					<div class="statsCell">
						<div>Accuracy:</div>
						<div>{item.accuracy}</div>
					</div>
					{/if}
					{#if item.charisma > 0}
					<div class="statsCell">
						<div>Charisma:</div>
						<div>{item.charisma}</div>
					</div>
					{/if}
				</div>
				{#if item.description != undefined && item.description != ""}
				<div class="cellHolder">
					<div>{item.description}</div>
				</div>
				{/if}
				{#if item.ability != undefined && item.ability != ""}
				<div class="cellHolder">
					<div>{item.ability}</div>
				</div>
				{/if}
			</div>
		</div>
	{/if}
{/key}

<style>
	a {
		text-decoration: none;
		color: #2d2d2d;
		font-weight: 400;
	}
	.item {
		/* display: flex; */
		flex-direction: column;
		width: 20rem;
		/* justify-content: space-between; */
		/* padding: 1rem; */
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
	.cursor {
		cursor: pointer;
	}
	.name {
		text-overflow: clip;
		overflow: hidden;
	}
	#open:checked + div > div > .name {
		text-overflow: initial;
		overflow-wrap: break-word;
		overflow: visible;
	}
</style>
