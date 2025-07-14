import { defineConfigs } from 'v-network-graph';
import { ForceLayout, type ForceNodeDatum, type ForceEdgeDatum } from 'v-network-graph/lib/force-layout';

const configs = defineConfigs({
    view: {
        grid: {
            visible: true,
            interval: 10,
            thickIncrements: 5,
            line: {
                color: "#e0e0e0",
                width: 1,
                dasharray: 1,
            },
            thick: {
                color: "#cccccc",
                width: 1,
                dasharray: 0,
            },
        },
        layoutHandler: new ForceLayout({
            positionFixedByDrag: false,
            positionFixedByClickWithAltKey: true,
            createSimulation: (d3, nodes, edges) => {
                // d3-force parameters
                const forceLink = d3.forceLink<ForceNodeDatum, ForceEdgeDatum>(edges).id((d: { id: any; }) => d.id)
                return d3
                    .forceSimulation(nodes)
                    .force("edge", forceLink.distance(50).strength(0.5))
                    .force("charge", d3.forceManyBody().strength(-25))
                    .alphaMin(0.001)
            }
        }),
        minZoomLevel: 0.1,
        maxZoomLevel: 16,
    },
    node: {
        normal: {
            color: "#ff6699",
        },
        hover: {
            color: "#ff99cc"
        },
        label: {
            visible: true,
            direction: "south",
            directionAutoAdjustment: true,
            color: "blue"
        },
    },
    edge: {
        gap: 50,
        normal: {
            color: "#ffb3cc",
        },
        hover: {
            color: "#ff99bb",
        },
        label: {
            color: "blue"
        },
        type: "curve",
    },
});

export default configs