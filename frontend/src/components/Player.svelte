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

	let modalUpdate = true;
	let showModal = false;
	export let id: number = 0;
</script>
{#if modalUpdate}
<Modal bind:showModal>
	<h2>Enter new HP and MANA</h2>
	<div>
		<label for="hp">HP</label>
		<input type="number" id="hp" name="hp" bind:value={hp}>
		<label for="mana">MANA</label>
		<input type="number" id="mana" name="mana" bind:value={mana}>
	</div>
	<button on:click={()=> {UpdateHPMANA()}}>Update</button>
</Modal>
{:else}
<Modal bind:showModal>
	<h2>Delete player?</h2>
	<button on:click={()=> {DeletePlayer()}}>DELETE</button>
</Modal>
{/if}
{#if p != undefined}
	<div class="player">
		<div class="info">
			<div class="type">Player</div>
			<div>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-missing-attribute -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<a on:click={()=>{modalUpdate=true;showModal=true;}}><span style="color:red;">HP</span><span style="color:blue">MANA</span></a>
				<a href={window.location.origin+"/players/update/"+id}>⚙️</a>
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<!-- svelte-ignore a11y-no-static-element-interactions -->
				<!-- svelte-ignore a11y-missing-attribute -->
				<a on:click={()=>{modalUpdate=false;showModal=true}}>❌</a>
			</div>
		</div>
		<div class="cellHolder">
			<div>{p.name}</div>
			<div>LVL {p.level}</div>
		</div>
		<div class="cellHolder">
			<div>{p.class}</div>
			<div>{p.race}</div>
			<div>{p.subrace}</div>
		</div>
		<div class="cellHolder">
			<div>HP {p.health}</div>
			<div>MANA {p.mana}</div>
		</div>
		<div class="cellHolder stats">
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
		<div class="cellHolder">
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
