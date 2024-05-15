<script lang="ts">
	import { onMount } from 'svelte';
	import Modal from './Modal.svelte'
	import SmallItem from './SmallItem.svelte';
	import { writable } from 'svelte/store';
  	import { HOST } from "$lib/host";

	let p: Player;
	let items = writable([])
	onMount(async function () {
		const res = await fetch(HOST + 'player/' + id);
		const data = await res.json();
		p = {
			id: data.id,
			name: data.name,
			level: data.level,
			health: data.health,
			mana: data.mana,
			class: data.class,
			race: data.race,
			subrace: data.subrace,
			strength: data.strength,
			endurance: data.endurance,
			perception: data.perception,
			intelligence: data.intelligence,
			agility: data.agility,
			accuracy: data.accuracy,
			charisma: data.charisma,
		};
		items.set(data.items)
	});

	let hp: number;
	let mana: number;
	async function UpdateHPMANA() {
		let data = {
			hp: hp,
			mana: mana,
		}
		fetch(HOST + "player/hpmana/" + p.id, {
        method: "PUT",
        body: JSON.stringify(data)
      })
      window.location.href = window.location.origin + "/players"
	}

	async function DeletePlayer() {
		fetch(HOST + "player/" + p.id, {
        method: "DELETE",
		})
		window.location.href = window.location.origin + "/players"
	}

	let showModal: boolean;
	let modalUpdate: boolean
	export let id: number = 0;
</script>
{#if modalUpdate}
<Modal bind:showModal>
	<div class="flex flex-col justify-center items-center">
	<h2>Enter new HP and MANA</h2>
	<div class="flex flex-col py-2">
		<label for="hp">HP</label>
		<input type="number" id="hp" name="hp" class="input input-sm input-bordered w-20" bind:value={hp}>
		<label for="mana">MANA</label>
		<input type="number" id="mana" name="mana" class="input input-sm input-bordered w-20" bind:value={mana}>
	</div>
	<button class="btn" on:click={()=> {UpdateHPMANA()}}>Update</button>
</div>
</Modal>
{:else}
<Modal bind:showModal>
	<div class="flex flex-col justify-center items-center">
	<h2>Delete player?</h2>
	<button class="btn" on:click={()=> {DeletePlayer()}}>DELETE</button>
</div>
</Modal>
{/if}
{#if p != undefined}
	<div class="player">
		<div class="info">
			<div>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<span on:click={()=>{modalUpdate=true;showModal = true}}><span class="text-red-500">HP</span><span class="text-blue-600">MANA</span></span>
			</div>
			<div>
				<a href={window.location.origin+"/players/update/"+id}>⚙️</a>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<!-- svelte-ignore a11y-missing-attribute -->
				<a on:click={()=>{modalUpdate=false;showModal=true}}>❌</a>
			</div>
		</div>
		<div class="cellHolder">
			<div class="underline underline-offset-1">{p.name}</div>
			<div class="badge badge-ghost badge-outline">LVL {p.level}</div>
		</div>
		<div class="cellHolder">
			<div class="badge badge-primary badge-outline">{p.class}</div>
			<div class="badge badge-secondary badge-outline">{p.race}</div>
			<div class="badge badge-accent badge-outline">{p.subrace}</div>
		</div>
		<div class="cellHolder">
			<div class="badge badge-outline badge-error">HP {p.health}</div>
			<div class="badge badge-outline badge-info">MANA {p.mana}</div>
		</div>
		<div class="cellHolder flex-col">
			<div class="statsCell">
				<div>Strength:</div>
				<div>{p.strength}</div>
			</div>
			<div class="statsCell">
				<div>Endurance:</div>
				<div>{p.endurance}</div>
			</div>
			<div class="statsCell">
				<div>Perception:</div>
				<div>{p.perception}</div>
			</div>
			<div class="statsCell">
				<div>Intelligence:</div>
				<div>{p.intelligence}</div>
			</div>
			<div class="statsCell">
				<div>Agility:</div>
				<div>{p.agility}</div>
			</div>
			<div class="statsCell">
				<div>Accuracy:</div>
				<div>{p.accuracy}</div>
			</div>
			<div class="statsCell">
				<div>Charisma:</div>
				<div>{p.charisma}</div>
			</div>
		</div>
		<div class="flex flex-wrap gap-4">
			{#if $items != null}
			{#each $items as item}
				<SmallItem id={item}/>
			{/each}
			{/if}
		</div>
	</div>
{/if}

<style>
	.player {
		display: flex;
		flex-direction: column;
		width: 20rem;
		justify-content: space-between;
		padding: 1rem;
		border: 3px solid;
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
