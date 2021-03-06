package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************
BRs: 7.1.2.2a certificatePolicies
This extension MUST be present and SHOULD NOT be marked critical.
************************************************/

import (
	"github.com/studyzy/fabric-sdk-go/third_party/github.com/zmap/zcrypto/x509"
	"github.com/studyzy/fabric-sdk-go/third_party/github.com/zmap/zlint/util"
)

type subCACertPolicyCrit struct{}

func (l *subCACertPolicyCrit) Initialize() error {
	return nil
}

func (l *subCACertPolicyCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *subCACertPolicyCrit) Execute(c *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(c, util.CertPolicyOID); e.Critical {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}

}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_certificate_policies_marked_critical",
		Description:   "Subordinate CA certificates certificatePolicies extension should not be marked as critical",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCACertPolicyCrit{},
	})
}
