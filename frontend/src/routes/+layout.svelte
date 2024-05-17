<script>
	import { HOST } from "$lib/host";
    import "tailwindcss/tailwind.css";

    let newdayKey = {}
    async function NewDay() {
      await fetch(HOST + "alcohol", {
        method: "POST",
      })
      newdayKey = {}
    }
  </script>

<div class="navbar bg-base-200 shadow-lg shadow-gray-700 sticky top-0">
  <div class="navbar-start" />
  <div class="navbar-center">
    <ul class="menu menu-horizontal px-1">
      <div class="join mr-4">
        <button id="menu-players" class="btn btn-ghost btn-outline join-item btn-sm rounded-l-full hover:btn-primary"><a href="/players">Players</a></button>
        <button id="menu-players-create" class="btn btn-outline btn-secondary join-item btn-sm rounded-r-full"><a href="/players/create">Create</a></button>
      </div>
      <div class="join mr-4">
        <button id="menu-items"class="btn btn-ghost btn-outline join-item btn-sm rounded-l-full hover:btn-primary"><a href="/items">Items</button>
        <button id="menu-items-create" class="btn btn-outline btn-secondary join-item btn-sm rounded-r-full"><a href="/items/create">Create</a></button>
      </div>
      <button id="menu-bag" class="btn btn-primary btn-outline join-item btn-sm rounded-full"><a href="/bag">Bag of Holding</a></button>
    </ul>
  </div>
  <div class="navbar-end mr-8 gap-8">
    <button on:click={NewDay} class="btn btn-outline btn-sm rounded-full">New Day</button>
    <label class="cursor-pointer grid place-items-center">
      <input type="checkbox" value="light" class="toggle theme-controller bg-base-content row-start-1 col-start-1 col-span-2"/>
      <svg class="col-start-1 row-start-1 stroke-base-100 fill-base-100" xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg>
      <svg class="col-start-2 row-start-1 stroke-base-100 fill-base-100" xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"/><path d="M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4"/></svg>
    </label>
  </div>
</div>
{#key newdayKey}
<div class="mx-10 py-10">
  <slot />
</div>
{/key}

<style>
  @import url('https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');
  * {
    font-family: Poppins;
  }
</style>