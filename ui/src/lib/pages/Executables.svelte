<script>
    import testIcon from '../../assets/test-img.png'
    import rightArrow from '../../assets/right-arrow.svg'
    import { navigate } from 'svelte-routing';

    let scrollY = 0;
    let maxY = 0;
    let screenHeight = 0;

    let bottomOpacity = 1;
    let topOpacity = 0;

    $: bottomOpacity = Math.max(scrollY+screenHeight-maxY, -20);
    $: topOpacity = Math.min(scrollY, 20);

    export let exeList = [
        "Google",
        "Steam",
        "Virus",
        "Trans Siberian Railway Simulator",
        "808",
        "Ts"
    ];
</script>

<svelte:window bind:scrollY bind:innerHeight={screenHeight} />


<div bind:clientHeight={maxY}>
    <div class="fixed bottom-0 left-0 w-full h-32 pointer-events-none" style="background: linear-gradient(180deg, rgba(255, 255, 255, 0.00) 30%, #FFFFFFF0 100%); opacity: {bottomOpacity/-20};"></div>
    <div class="fixed top-0 left-0 w-full h-32 pointer-events-none" style="background: linear-gradient(0deg, rgba(255, 255, 255, 0.00) 50%, #FFFFFFE0 100%); opacity: {topOpacity/20};"></div>
    {#each exeList as exe}
        <div class="flex py-3 px-5 h-24 w-full gap-5 items-center" style="border-bottom: 2px solid #E1E1E1;">
            <div class="min-w-8 max-w-8"><img class="w-full" src={testIcon} alt="icon"/></div>
            <div class="grow overflow-hidden">
                <div class="font-normal text-2xl text-ellipsis overflow-hidden text-nowrap">{exe}</div>
                <div class="font-normal text-base text-[#00000080] text-ellipsis overflow-hidden text-nowrap">Signed by Gus Inc.</div>
            </div>
            <button on:click={() => navigate("/exe-details")} >
                <img src={rightArrow} alt="select" class="nextButton w-12 rounded-[5px] transition-all" />
            </button>
        </div>
    {/each}
</div>

<style>
.nextButton:hover {
    background-color: #d9d9d9;
}
</style>