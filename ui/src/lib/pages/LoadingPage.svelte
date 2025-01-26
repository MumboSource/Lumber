<script>
    import { navigate } from "svelte-routing";

    setTimeout(() => {
        navigate("/exe-list") 
    }, 1000)
</script>

<div class="flex flex-col w-screen h-screen items-center justify-center gap-6">
    <img alt="logo" />
    <div class="flex flex-col items-center">
        Lumber is loading...
        <img alt="loading gif" />
    </div>
</div>