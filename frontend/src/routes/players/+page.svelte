<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
  	import { HOST } from "$lib/host";
	import Player from '../../components/Player.svelte';

	let players = writable([]);
	onMount(async () => {
		const res = await fetch(HOST + 'players');
		const data = await res.json();
		players.set(data);
	});

</script>

<div style="display: flex;">
	<h1>Players</h1>
	<button class="btn" on:click={()=>{window.location.href = window.location.origin+"/players/create"}}>Create Player</button>
</div>
<div class="PlayerHolder">
	{#if $players != null}
	{#each $players as player}
		<Player id={player} />
	{/each}
	{:else}
	<h1>No players</h1>
	{/if}
</div>

<style>
	.PlayerHolder {
		display: flex;
        flex-wrap: wrap;
		gap: 10px;
	}

	.btn {
		margin: auto;
	}
</style>
