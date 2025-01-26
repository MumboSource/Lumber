<script>
    import { navigate } from "svelte-routing";

    setTimeout(() => {
        navigate("/exe-list") 
    }, 2500)
</script>

<div class="flex flex-col w-screen h-screen items-center justify-center gap-6">
    <img src="/icon.png" class="w-24 h-24"/>
    <div class="flex flex-col items-center">
        Lumber is loading...
        <img src="/spinner.png" class="w-10 h-10 animate-spin" />
    </div>
</div>