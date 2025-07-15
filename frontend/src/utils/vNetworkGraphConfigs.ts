import { defineConfigs } from "v-network-graph";
import {
  ForceLayout,
  type ForceNodeDatum,
  type ForceEdgeDatum,
} from "v-network-graph/lib/force-layout";

export function createGraphConfig(
  theme: "light" | "dark",
  nodeSelectable: boolean | number = false
) {
  const isDark = theme == "dark";

  return defineConfigs({
    view: {
      grid: {
        visible: true,
        interval: 10,
        thickIncrements: 5,
        line: {
          color: isDark ? "#444" : "#e0e0e0",
          width: 1,
          dasharray: 1,
        },
        thick: {
          color: isDark ? "#666" : "#ccc",
          width: 1,
          dasharray: 0,
        },
      },
      layoutHandler: new ForceLayout({
        positionFixedByDrag: false,
        positionFixedByClickWithAltKey: true,
        createSimulation: (d3, nodes, edges) => {
          const forceLink = d3
            .forceLink<ForceNodeDatum, ForceEdgeDatum>(edges)
            .id((d: { id: any }) => d.id);
          return d3
            .forceSimulation(nodes)
            .force("edge", forceLink.distance(60).strength(0.5))
            .force("charge", d3.forceManyBody().strength(-0.05))
            .alphaMin(0.001);
        },
      }),
      minZoomLevel: 0.5,
      maxZoomLevel: 10,
    },
    node: {
      selectable: nodeSelectable ? 2 : false,
      normal: {
        color: isDark ? "#ff6699" : "#d13b6f",
      },
      hover: {
        color: isDark ? "#ff99cc" : "#f273a3",
      },
      label: {
        visible: true,
        direction: "south",
        directionAutoAdjustment: true,
        color: isDark ? "white" : "black",
      },
    },
    edge: {
      selectable: false,
      gap: 50,
      normal: {
        color: isDark ? "#ffb3cc" : "#f18ca3",
      },
      hover: {
        color: isDark ? "#ff99bb" : "#ff6f96",
      },
      label: {
        color: isDark ? "white" : "black",
      },
      type: "curve",
    },
  });
}
