<script lang="ts">
    import type client from "$lib/util/api/client";

    export let group: client.group.types.GroupData;
</script>


<h3>{group.name}</h3>
<p>{group.description}</p>