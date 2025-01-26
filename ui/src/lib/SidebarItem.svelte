<script>
    import { Link } from "svelte-routing";

    export let url;
    export let text;
    export let iconName;
    export let selected = false;
</script>

<Link to={url}>
    <div class="flex flex-row text-xl gap-2 px-3 py-4 {selected?"border-blue-500 border-l-4 bg-[#F6F6F6] font-bold cursor-default":"hover:bg-[#F6F6F6]"}">
        <img src="/{iconName}{selected?"-active":""}.svg" alt="" class="w-8 h-8">
        <p>{text}</p>
    </div>
</Link>