permit (
    principal is CG,
    action,
    resource == Capability::"wallet"
)
when {
    context.crId in principal.cr &&
    context.CRCGActionSet.contains(action) &&
    resource in context.CRCGResourceSet
};

