<script lang="ts">
    import CircleExclamation from "../../../components/svg/CircleExclamation.svelte";

    export let formQuestion:string
    export let questionNum:number
</script>

<div class="flex justify-center items-start md:items-center md:h-auto md:mt-8">
    <div class="flex justify-center">
        <div class="flex flex-grow flex-col md:flex-row md:justify-center items-center md:w-4/5 gap-4">
            <h1 class="text-xl text-primary font-bold text-center">Spørsmål {Number(questionNum) + 1}</h1>
            <div class="border-b-2 md:border-r-2 border-primary md:min-h-24 w-full md:w-0 md:h-full" />
            <p class="text-sm">{formQuestion}</p>
        </div>
    </div>
    <div class="flex justify-end md:pr-20 items-center md:w-1/3">
        <a class="-" href="/reportbug">
            <CircleExclamation width="16" height="16" />    
        </a>
    </div>
</div>