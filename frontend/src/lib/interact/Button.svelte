<script lang="ts">
    export let span : number = 2;
    export let bg: string = "var(--COLOR-2)";
    export let color: string = "var(--COLOR-4)";
</script>


<button style={`grid-column: span ${span}; background-color: ${bg}; color:${color};`} on:click>
    <slot/>
</button>

<style>
    button {
        cursor: pointer;
        padding: calc(var(--SPACING) * 0.6);
        margin: 0px;
        /* background-color: var(--COLOR-2); */
        outline: none;
        border: none;
        /* color: inherit; */
        text-transform: uppercase;
        text-decoration: none;
        font-size: 1rem;
        box-shadow: inset 0px 0px 0px -3px var(--COLOR-3);
        transition: all .1s ease;
    }
    /* button:hover, button:active, button:focus  {
        box-shadow: inset 0px -5px 0px -3px var(--COLOR-3);
    } */
</style>