<script>
    import min from '../assets/min.svg'
    import max from '../assets/max.svg'
    import unMax from '../assets/fullscreen-exit.svg'
    import close from '../assets/close.svg'
    import { navigate } from 'svelte-routing'

    import testIcon from '../assets/test-img.png'
    import backArrow from '../assets/left-arrow.svg'
    import { onMount } from 'svelte';

    let maximized = false;

    export let maximizable = false;
    export let title = "";
    export let icon = testIcon;
    export let backUrl = "";

    onMount(async () => {
        maximized = await window.electron.isMaximized();
    })
</script>

<div class="flex h-16 w-full pl-3 gap-3 items-center" style="border-bottom: 2px solid #E1E1E1;">
    <div class="min-w-8 max-w-8 py-3" style="app-region: drag;"><img class="w-full" src={icon} alt="icon"/></div>
    {#if backUrl != ""}
        <button on:click={() => navigate(backUrl)}>
            <img src={backArrow} alt="select" class="nextButton w-12 rounded-[5px] transition-all" />
        </button>
    {/if}
    <div class="text-2xl grow h-full overflow-hidden text-nowrap text-ellipsis flex items-center" style="app-region: drag;">{title}</div>
    

    <div class="flex h-full items-center justify-around" style="border-left: 2px solid #E1E1E1;">
        <button class="nextButton h-full px-3" on:click={() => window.electron.minimize()} >
            <img src={min} alt="select" class="scale-75 w-12 transition-all" />
        </button>
        {#if maximizable}
            {#if !maximized}
            <button class="nextButton h-full px-3" on:click={() => {window.electron.maximize(); maximized = true}} >
                <img src={max} alt="select" class="scale-75 w-12 transition-all" />
            </button>
            {:else}
            <button class="nextButton h-full px-3" on:click={() => {window.electron.unmaximize(); maximized = false}} >
                <img src={unMax} alt="select" class="scale-75 w-12 transition-all" />
            </button>
            {/if}
        {/if}
        <button class="nextButton h-full px-3" on:click={() => window.electron.close()} >
            <img src={close} alt="select" class="scale-75 w-12 transition-all" />
        </button>
    </div>
</div>

<style>
    .nextButton:hover {
        background-color: #d9d9d9;
    }
</style>