<script>
	import { change_water_by, get_water } from '$lib/api/game';
	let water = $state(get_water());
	let water_change = $state(10);

	function reload() {
		water = get_water();
	}

    function change() {
        change_water_by(water_change)
        reload()
    }
</script>

<div class="p-4 backdrop:blur-3xl bg-zinc-800 size-min text-white">
	<h1 class="py-2 text-xl font-bold">Water Amount</h1>
	<form onsubmit={change} class="flex flex-row gap-2">
		<input type="number" class="bg-zinc-700 rounded-md px-4 py-2" bind:value={water_change} />
		<button type="submit">Change</button>
	</form>
	<form onsubmit={reload} class="flex flex-row gap-2">
        <button type="submit" >Reload</button>
        {#await water}
            <p>Loading...</p>
        {:then data}
            <p>{data.water}</p>
        {:catch err}
            Error: {err}
        {/await}
    </form>
</div>
