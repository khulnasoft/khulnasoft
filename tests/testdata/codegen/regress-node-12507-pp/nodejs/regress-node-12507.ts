import * as khulnasoft from "@khulnasoft/khulnasoft";
import * as aws from "@khulnasoft/aws";

const config = new khulnasoft.Config();
const localGatewayVirtualInterfaceGroupId = config.require("localGatewayVirtualInterfaceGroupId");
const rts = aws.ec2.getLocalGatewayRouteTablesOutput({
    filters: [{
        name: "tag:kubernetes.io/kops/role",
        values: ["private*"],
    }],
});
const routes: aws.ec2.LocalGatewayRoute[] = [];
rts.ids.length.apply(rangeBody => {
    for (const range = {value: 0}; range.value < rangeBody; range.value++) {
        routes.push(new aws.ec2.LocalGatewayRoute(`routes-${range.value}`, {
            destinationCidrBlock: "10.0.1.0/22",
            localGatewayRouteTableId: rts.apply(rts => rts.ids[range.value]),
            localGatewayVirtualInterfaceGroupId: localGatewayVirtualInterfaceGroupId,
        }));
    }
});
