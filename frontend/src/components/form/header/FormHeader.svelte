<script lang="ts">
    import CircleExclamation from "../../../components/svg/CircleExclamation.svelte";

    export let formQuestion:string
    export let questionNum:number
</script>

<div class="flex justify-center items-center h-20">
    <div class="w-1/3"/>
    <div class="flex justify-center items-center w-2/3 gap-4">
        <h1 class="text-xl text-primary font-bold text-center">Spørsmål {Number(questionNum) + 1}</h1>
        <div class="border-r-2 border-primary h-24"></div>
        <p class="text-sm">{formQuestion}</p>
    </div>
    <div class="flex justify-end pr-20 items-center w-1/3">
        <a class="-" href="/errorreport">
            <CircleExclamation width="16" height="16" />    
        </a>
    </div>
</div>