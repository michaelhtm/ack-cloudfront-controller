    // We need to do this manually because the ETag field (required for the
    // IfMatch field when updating or deleting a resource in CloudFront) is
    // outside the wrapper field of `Distribution` in the response and
    // therefore not picked up by SetResource code generation.
	if resp.ETag != nil {
		ko.Status.ETag = resp.ETag
	} 

	if ko.Spec.Tags != nil {
		ackcondition.SetSynced(&resource{ko}, corev1.ConditionFalse, nil, nil)
	}