<script>
	import ButtonComponent from "../components/userform/inputs/ButtonComponent.svelte";
</script>
<div class="flex flex-col justify-between items-center h-full">
    <div class="flex h-full items-center">
        <div class="flex flex-col w-3/5 gap-4 px-32">
            <h1 class="text-7xl text-primary font-bold">Helseveileder</h1>
            <p>Vi utfører en undersøkelse om kunstig intelligens (AI) i medisin. Vi ønsker å forstå hvordan AI kan fungere som en copilot for leger og pasienter.</p>
            <p>Undersøkelsen er helt anonym.</p>
            <ul>
                <p>I denne undersøkelsen vil vi at du skal vurdere empati, kunnskap og hjelpsomhet i ulike svar på medisinske spørsmål.</p>
                <li><strong>Empati:</strong> Evnen til å oppdage og erkjenne andres følelser og tanker.</li>
                <li><strong>Kunnskap:</strong> Nøyaktighet og relevans av den medisinske informasjonen gitt.</li>
                <li><strong>Hjelpsomhet:</strong> Evnen til å gi nyttig og praktisk råd eller støtte.</li>
            </ul>
        </div>
        <div class="flex justify-center items-center w-2/5">
            <a href="/userform">
                <ButtonComponent text="Start undersøkelse" filled={false} />
            </a>
        </div>
    </div>
    <div class="flex justify-center items-center gap-8 text-primary font-bold">
        <button><a href="/personvern">Personvern</a></button>
        <button><a href="/info">Mer info</a></button>
    </div>
</div>

<style lang="postcss">
    :root {
        color: theme(colors.content);
        
    }
    li {
        list-style-type: disc;
        margin-left: 2rem;
        font-size: large;
    }
    p {
        font-size: large;
    }
    button {
        font-size: large;
    }
</style>