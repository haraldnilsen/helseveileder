<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import CheckMark from '../../svg/CheckMark.svelte';
    const dispatch = createEventDispatcher();
    function handleButtonClick(selectedValue: string) {
        dispatch('update', selectedValue);
    }
    export let options:string[]
    export let selected: string
</script>

<div class="flex justify-between w-full md:pl-5">
{#each options as data, index (data)}
    <button name={data} on:click={() => handleButtonClick(data)} class="flex flex-col justify-start items-center w-12 text-center hover:cursor-pointer">
        <div class={`${selected == data && "bg-primary"} flex justify-center items-center h-6 w-6 rounded-full border-2 border-primary hover:bg-primary`}>
            {#if selected == data}
                <CheckMark width=10/>
            {/if}
        </div>
        <label class="text-primary text-sm mt-1 hover:cursor-pointer" for={data}>{data}</label>
    </button>
{/each}
</div>
